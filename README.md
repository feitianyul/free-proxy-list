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

最后更新：2026-05-16 20:00:05 UTC（2026-05-17 04:00:05 UTC+8）

**代理总数：50**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1651ms | ✓ 1176ms | ✓ 1482ms | ✓ 1346ms | ✓ 1144ms | http |
| 212.58.132.5:8888 | ✓ 1299ms | 否 | ✓ 1420ms | ✓ 1445ms | ✓ 1115ms | http |
| 185.200.188.234:10001 | ✓ 1674ms | 否 | ✓ 1460ms | 否 | ✓ 1836ms | http |
| 113.160.132.26:8080 | ✓ 1961ms | ✓ 1998ms | ✓ 1154ms | ✓ 1615ms | ✓ 1277ms | http |
| 218.108.131.186:17890 | ✓ 1867ms | ✓ 1064ms | ✓ 847ms | ✓ 1076ms | ✓ 932ms | http |
| 170.106.136.181:31002 | 否 | ✓ 1245ms | ✓ 1007ms | 否 | ✓ 1021ms | http |
| 42.114.172.179:2075 | ✓ 1905ms | 否 | ✓ 1774ms | 否 | ✓ 1825ms | http |
| 84.47.150.125:1080 | ✓ 1237ms | 否 | ✓ 1670ms | 否 | ✓ 1690ms | http |
| 150.107.140.238:3128 | ✓ 1796ms | 否 | ✓ 1091ms | 否 | ✓ 1167ms | http |
| 91.242.229.129:8092 | 否 | ✓ 1989ms | 否 | ✓ 1618ms | ✓ 1219ms | http |
| 157.0.142.246:10057 | 否 | ✓ 1262ms | ✓ 1022ms | ✓ 1326ms | 否 | http |
| 114.214.170.41:27890 | 否 | 否 | ✓ 1313ms | ✓ 1462ms | ✓ 1191ms | http |
| 103.21.220.141:3128 | ✓ 883ms | 否 | ✓ 846ms | ✓ 1079ms | ✓ 1126ms | http |
| 128.199.114.189:9090 | ✓ 1122ms | 否 | ✓ 1135ms | ✓ 1688ms | ✓ 1330ms | http |
| 148.230.4.241:999 | ✓ 781ms | 否 | ✓ 868ms | ✓ 1647ms | ✓ 1426ms | http |
| 45.129.141.143:3128 | ✓ 837ms | ✓ 1723ms | ✓ 1501ms | 否 | ✓ 1665ms | http |
| 137.59.47.73:3128 | ✓ 1047ms | ✓ 1535ms | ✓ 1213ms | 否 | ✓ 1391ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1571ms | ✓ 1362ms | ✓ 1638ms | ✓ 1322ms | http |
| 8.154.21.175:3128 | ✓ 1080ms | ✓ 1420ms | ✓ 1005ms | ✓ 1382ms | ✓ 1167ms | http |
| 121.130.199.80:24007 | 否 | ✓ 1811ms | ✓ 1795ms | ✓ 1530ms | ✓ 1374ms | http |
| 47.105.98.23:3128 | ✓ 1650ms | 否 | ✓ 1312ms | 否 | ✓ 1882ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1467ms | ✓ 1314ms | 否 | ✓ 1067ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1491ms | ✓ 1208ms | ✓ 1061ms | http |
| 45.125.67.37:8443 | ✓ 1819ms | 否 | ✓ 1792ms | ✓ 1900ms | 否 | http |
| 173.212.246.157:3128 | ✓ 947ms | 否 | ✓ 1329ms | ✓ 1842ms | ✓ 1397ms | http |
| 2.27.32.81:3128 | ✓ 1354ms | ✓ 1978ms | ✓ 1763ms | 否 | 否 | http |
| 107.175.85.198:1080 | ✓ 1846ms | ✓ 1362ms | ✓ 1066ms | 否 | 否 | http |
| 42.114.172.179:2045 | ✓ 1764ms | 否 | ✓ 1863ms | 否 | ✓ 1813ms | http |
| 166.88.55.83:7890 | ✓ 821ms | ✓ 1348ms | ✓ 953ms | ✓ 1055ms | ✓ 867ms | http |
| 121.230.8.109:1080 | ✓ 1343ms | ✓ 1534ms | ✓ 1268ms | 否 | ✓ 1415ms | http |
| 210.76.192.50:10808 | ✓ 1105ms | ✓ 1328ms | ✓ 1419ms | ✓ 1418ms | 否 | http |
| 86.104.72.220:1081 | ✓ 440ms | ✓ 913ms | ✓ 294ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 755ms | 否 | ✓ 1929ms | 否 | ✓ 1854ms | http |
| 5.129.248.58:3128 | ✓ 1201ms | ✓ 1898ms | 否 | 否 | ✓ 1864ms | http |
| 128.199.113.85:9090 | ✓ 945ms | 否 | ✓ 1203ms | ✓ 1362ms | ✓ 1064ms | http |
| 3.15.187.17:1080 | ✓ 843ms | ✓ 1516ms | 否 | ✓ 1837ms | 否 | http |
| 103.147.152.12:1080 | ✓ 803ms | ✓ 1407ms | ✓ 1789ms | ✓ 1442ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1848ms | 否 | ✓ 1699ms | 否 | ✓ 1693ms | http |
| 159.89.31.62:8080 | ✓ 1029ms | ✓ 1595ms | ✓ 1609ms | ✓ 1616ms | 否 | http |
| 3.101.133.120:80 | ✓ 677ms | ✓ 1518ms | ✓ 1467ms | ✓ 1299ms | ✓ 1073ms | http |
| 193.181.35.217:8118 | ✓ 1044ms | ✓ 1595ms | ✓ 1199ms | ✓ 1513ms | ✓ 1493ms | http |
| 45.88.0.111:3128 | ✓ 506ms | 否 | ✓ 809ms | ✓ 1640ms | 否 | http |
| 190.12.150.244:999 | ✓ 879ms | ✓ 1513ms | ✓ 1079ms | 否 | 否 | http |
| 38.211.245.131:999 | ✓ 1573ms | 否 | ✓ 1005ms | 否 | ✓ 1956ms | http |
| 43.156.90.221:10808 | ✓ 1584ms | ✓ 1309ms | ✓ 952ms | ✓ 1231ms | ✓ 996ms | http |
| 45.153.231.229:8080 | ✓ 971ms | 否 | ✓ 1980ms | ✓ 1992ms | ✓ 1984ms | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1851ms | ✓ 1726ms | ✓ 1211ms | http |
| 121.230.8.144:1080 | ✓ 1274ms | ✓ 1626ms | ✓ 1440ms | ✓ 1849ms | ✓ 1301ms | http |
| 180.103.19.207:1080 | 否 | ✓ 1621ms | ✓ 1254ms | ✓ 1608ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1565ms | ✓ 1403ms | ✓ 1404ms | ✓ 1943ms | 否 | http |

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
