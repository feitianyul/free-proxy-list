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

最后更新：2026-04-04 12:31:48 UTC（2026-04-04 20:31:48 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 752ms | ✓ 1348ms | ✓ 1138ms | ✓ 1222ms | ✓ 1159ms | http |
| 1.231.81.166:3128 | ✓ 992ms | ✓ 1378ms | 否 | ✓ 1209ms | ✓ 1002ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1807ms | ✓ 1056ms | ✓ 1448ms | ✓ 1086ms | http |
| 95.213.217.168:52004 | ✓ 1231ms | ✓ 1636ms | 否 | 否 | ✓ 1760ms | http |
| 167.103.115.102:8800 | ✓ 1247ms | 否 | ✓ 1349ms | ✓ 1454ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1602ms | 否 | ✓ 1747ms | ✓ 1810ms | ✓ 1689ms | http |
| 37.204.248.33:443 | ✓ 1473ms | 否 | ✓ 1691ms | ✓ 1967ms | 否 | http |
| 45.140.147.82:1081 | ✓ 649ms | ✓ 1264ms | ✓ 1017ms | ✓ 1596ms | ✓ 1173ms | http |
| 159.223.71.162:8080 | ✓ 865ms | 否 | 否 | ✓ 1749ms | ✓ 984ms | http |
| 111.227.254.12:22222 | ✓ 1190ms | 否 | 否 | ✓ 1589ms | ✓ 1213ms | http |
| 167.103.144.127:8800 | ✓ 1504ms | 否 | ✓ 1213ms | ✓ 1439ms | ✓ 1305ms | http |
| 209.38.154.7:1080 | 否 | ✓ 904ms | 否 | ✓ 1876ms | ✓ 651ms | http |
| 167.103.31.122:8800 | ✓ 1520ms | 否 | ✓ 1269ms | 否 | ✓ 1493ms | http |
| 180.250.219.58:53281 | ✓ 1958ms | 否 | ✓ 1729ms | 否 | ✓ 1998ms | http |
| 8.219.97.248:80 | ✓ 1646ms | 否 | ✓ 1265ms | ✓ 1681ms | 否 | http |
| 35.225.22.61:80 | ✓ 172ms | ✓ 1203ms | ✓ 135ms | ✓ 1158ms | ✓ 796ms | http |
| 111.227.254.9:22222 | ✓ 1156ms | ✓ 1546ms | ✓ 1150ms | 否 | ✓ 1186ms | http |
| 147.161.239.240:8800 | ✓ 1014ms | ✓ 1740ms | ✓ 1246ms | ✓ 1521ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1071ms | 否 | ✓ 1488ms | 否 | ✓ 1567ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1936ms | ✓ 1313ms | 否 | ✓ 1726ms | http |
| 45.167.125.21:999 | ✓ 588ms | ✓ 1843ms | ✓ 567ms | ✓ 1710ms | ✓ 1451ms | http |
| 38.34.178.152:8451 | 否 | 否 | ✓ 1637ms | ✓ 1503ms | ✓ 1928ms | http |
| 38.145.220.49:8444 | ✓ 1356ms | ✓ 1916ms | 否 | ✓ 1424ms | 否 | http |
| 38.34.179.26:8446 | ✓ 937ms | 否 | ✓ 944ms | 否 | ✓ 1978ms | http |
| 150.241.71.15:1080 | ✓ 1310ms | 否 | ✓ 1902ms | ✓ 1366ms | ✓ 1034ms | http |
| 38.34.179.6:8449 | 否 | ✓ 1226ms | ✓ 770ms | 否 | ✓ 806ms | http |
| 24.152.59.110:999 | ✓ 1375ms | 否 | ✓ 657ms | ✓ 1798ms | ✓ 1454ms | http |
| 45.140.147.82:1082 | 否 | ✓ 1731ms | ✓ 1219ms | ✓ 1613ms | 否 | http |
| 103.193.145.22:8082 | ✓ 1826ms | 否 | 否 | ✓ 1567ms | ✓ 1505ms | http |
| 103.191.254.1:7777 | ✓ 1818ms | 否 | ✓ 1853ms | 否 | ✓ 1564ms | http |
| 223.16.170.103:80 | ✓ 1377ms | 否 | ✓ 1874ms | ✓ 1236ms | 否 | http |
| 45.136.131.48:8446 | ✓ 1086ms | ✓ 912ms | ✓ 978ms | ✓ 1693ms | ✓ 981ms | http |
| 45.136.131.51:8446 | ✓ 1087ms | ✓ 985ms | ✓ 902ms | ✓ 1713ms | ✓ 969ms | http |
| 45.136.131.68:8446 | ✓ 1584ms | ✓ 834ms | ✓ 1127ms | 否 | ✓ 673ms | http |
| 38.145.220.32:8449 | ✓ 1105ms | ✓ 1847ms | ✓ 1644ms | ✓ 1910ms | ✓ 724ms | http |
| 38.34.179.29:8452 | 否 | 否 | ✓ 495ms | ✓ 1837ms | ✓ 1360ms | http |
| 168.222.254.26:8888 | ✓ 1089ms | ✓ 1788ms | ✓ 1501ms | 否 | 否 | http |
| 45.12.151.226:2829 | 否 | ✓ 1685ms | 否 | ✓ 1700ms | ✓ 1240ms | http |
| 59.46.216.131:30001 | ✓ 1884ms | 否 | ✓ 1200ms | 否 | ✓ 1199ms | http |
| 159.223.71.162:443 | ✓ 1383ms | 否 | 否 | ✓ 1196ms | ✓ 969ms | http |
| 54.222.174.194:80 | 否 | ✓ 1752ms | ✓ 1855ms | ✓ 1852ms | 否 | http |
| 45.136.131.67:8451 | ✓ 1535ms | ✓ 1082ms | ✓ 638ms | ✓ 1696ms | ✓ 858ms | http |
| 45.125.67.37:443 | ✓ 1737ms | 否 | ✓ 1035ms | ✓ 1503ms | ✓ 1787ms | http |
| 208.87.243.199:7878 | ✓ 1201ms | 否 | ✓ 1900ms | ✓ 1113ms | ✓ 1840ms | http |
| 16.78.119.130:443 | 否 | ✓ 1882ms | ✓ 1998ms | ✓ 1956ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1864ms | 否 | ✓ 1707ms | ✓ 1703ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1218ms | 否 | ✓ 1434ms | ✓ 1758ms | ✓ 1716ms | http |
| 5.104.87.17:8050 | 否 | 否 | ✓ 1645ms | ✓ 1650ms | ✓ 1220ms | http |
| 45.136.198.40:3128 | ✓ 1221ms | 否 | 否 | ✓ 1880ms | ✓ 1735ms | http |
| 45.129.141.143:3128 | ✓ 1204ms | ✓ 1695ms | 否 | ✓ 1879ms | 否 | http |
| 47.110.42.192:9003 | ✓ 1672ms | ✓ 1593ms | ✓ 1505ms | ✓ 1692ms | ✓ 1627ms | http |
| 205.164.46.6:3094 | ✓ 319ms | 否 | ✓ 1996ms | ✓ 1957ms | 否 | http |
| 223.16.170.103:3128 | ✓ 1018ms | 否 | ✓ 1191ms | ✓ 1281ms | ✓ 1276ms | http |
| 61.52.131.172:8443 | ✓ 1022ms | ✓ 1315ms | ✓ 1606ms | ✓ 1259ms | ✓ 1081ms | http |
| 106.117.208.101:7890 | ✓ 1104ms | ✓ 1431ms | 否 | ✓ 1800ms | ✓ 1490ms | http |

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
