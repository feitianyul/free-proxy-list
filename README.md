**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-13 14:47:12 UTC（2026-04-13 22:47:12 UTC+8）

**代理总数：48**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1543ms | ✓ 1903ms | ✓ 746ms | ✓ 1055ms | ✓ 835ms | http |
| 167.103.34.108:8800 | ✓ 1931ms | 否 | ✓ 1526ms | ✓ 1662ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1875ms | 否 | ✓ 1807ms | ✓ 1615ms | ✓ 1557ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1240ms | ✓ 1348ms | ✓ 1091ms | http |
| 43.156.132.113:3128 | ✓ 734ms | 否 | ✓ 675ms | ✓ 1522ms | ✓ 840ms | http |
| 167.103.115.102:8800 | ✓ 936ms | 否 | ✓ 900ms | ✓ 1277ms | ✓ 959ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1278ms | ✓ 1134ms | ✓ 849ms | http |
| 79.132.136.58:3128 | ✓ 1626ms | 否 | ✓ 1950ms | ✓ 1742ms | ✓ 1295ms | http |
| 114.237.77.253:1080 | 否 | ✓ 1279ms | ✓ 1959ms | 否 | ✓ 1625ms | http |
| 168.110.52.228:3128 | ✓ 1692ms | 否 | 否 | ✓ 1932ms | ✓ 1773ms | http |
| 59.46.216.131:30001 | ✓ 1040ms | ✓ 1271ms | 否 | ✓ 1937ms | 否 | http |
| 45.167.125.21:999 | ✓ 1824ms | 否 | ✓ 1552ms | ✓ 1936ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1177ms | 否 | ✓ 826ms | ✓ 1211ms | ✓ 1088ms | http |
| 167.103.31.122:8800 | ✓ 1291ms | 否 | ✓ 1389ms | 否 | ✓ 1591ms | http |
| 20.210.76.178:8561 | ✓ 598ms | ✓ 1355ms | ✓ 542ms | ✓ 782ms | ✓ 600ms | http |
| 20.210.76.175:8561 | ✓ 600ms | ✓ 1345ms | ✓ 579ms | ✓ 782ms | ✓ 600ms | http |
| 20.27.15.49:8561 | ✓ 590ms | ✓ 955ms | ✓ 547ms | ✓ 780ms | ✓ 594ms | http |
| 20.210.76.104:8561 | ✓ 625ms | ✓ 1260ms | ✓ 486ms | ✓ 785ms | ✓ 610ms | http |
| 147.161.239.240:8800 | ✓ 1401ms | 否 | ✓ 1504ms | ✓ 1854ms | ✓ 1683ms | http |
| 181.78.44.63:999 | ✓ 964ms | 否 | ✓ 1361ms | ✓ 1857ms | ✓ 1384ms | http |
| 95.214.9.93:3128 | ✓ 912ms | 否 | ✓ 1328ms | ✓ 1853ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1047ms | 否 | ✓ 1400ms | 否 | ✓ 1414ms | http |
| 46.30.46.133:3128 | ✓ 1139ms | 否 | ✓ 1637ms | ✓ 1900ms | 否 | http |
| 8.209.238.110:47701 | ✓ 536ms | 否 | ✓ 755ms | ✓ 810ms | ✓ 670ms | http |
| 120.92.108.86:7890 | ✓ 1408ms | 否 | 否 | ✓ 1964ms | ✓ 1728ms | http |
| 103.157.200.126:3128 | ✓ 1701ms | 否 | 否 | ✓ 1999ms | ✓ 1571ms | http |
| 159.223.225.118:8888 | ✓ 1669ms | 否 | 否 | ✓ 1858ms | ✓ 1409ms | http |
| 210.223.44.230:3128 | ✓ 807ms | ✓ 1129ms | ✓ 652ms | ✓ 850ms | ✓ 683ms | http |
| 36.103.198.235:7890 | ✓ 1100ms | ✓ 1346ms | ✓ 1284ms | 否 | 否 | http |
| 103.22.99.43:8085 | 否 | 否 | ✓ 1652ms | ✓ 1435ms | ✓ 1337ms | http |
| 103.157.117.116:8080 | ✓ 1857ms | 否 | ✓ 1761ms | 否 | ✓ 1658ms | http |
| 165.99.56.135:8080 | 否 | 否 | ✓ 1909ms | ✓ 1350ms | ✓ 1442ms | http |
| 139.198.113.42:10023 | ✓ 1938ms | ✓ 1530ms | ✓ 1039ms | ✓ 1096ms | ✓ 1224ms | http |
| 218.108.131.186:17890 | ✓ 764ms | ✓ 983ms | ✓ 807ms | ✓ 1039ms | ✓ 865ms | http |
| 140.227.61.201:3128 | ✓ 1525ms | 否 | ✓ 1781ms | ✓ 1820ms | 否 | http |
| 171.227.167.109:1005 | ✓ 929ms | 否 | ✓ 1291ms | ✓ 1238ms | ✓ 939ms | http |
| 103.166.158.201:3125 | ✓ 1794ms | 否 | ✓ 1460ms | ✓ 1617ms | ✓ 1541ms | http |
| 103.82.23.118:5234 | ✓ 1916ms | 否 | 否 | ✓ 1530ms | ✓ 1356ms | http |
| 101.43.127.100:8877 | ✓ 1539ms | ✓ 1671ms | ✓ 1365ms | 否 | ✓ 1883ms | http |
| 8.219.97.248:80 | ✓ 1078ms | 否 | 否 | ✓ 1119ms | ✓ 1178ms | http |
| 116.203.112.97:3128 | ✓ 692ms | 否 | ✓ 1534ms | 否 | ✓ 1664ms | http |
| 114.237.77.199:1080 | ✓ 1828ms | ✓ 1193ms | 否 | ✓ 1157ms | ✓ 1505ms | http |
| 180.103.19.123:1080 | ✓ 1088ms | 否 | ✓ 1337ms | ✓ 1993ms | ✓ 1805ms | http |
| 61.52.131.172:8443 | ✓ 953ms | ✓ 1152ms | ✓ 921ms | ✓ 1203ms | ✓ 910ms | http |
| 194.67.99.223:1080 | ✓ 886ms | 否 | ✓ 1883ms | 否 | ✓ 1486ms | http |
| 110.42.37.202:20005 | ✓ 1435ms | ✓ 1380ms | ✓ 1484ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1347ms | 否 | ✓ 1862ms | 否 | ✓ 1502ms | http |
| 181.119.84.80:999 | ✓ 1756ms | ✓ 1947ms | ✓ 1716ms | 否 | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
