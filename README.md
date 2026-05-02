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

最后更新：2026-05-02 22:39:41 UTC（2026-05-03 06:39:41 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.45.178.211:14658 | ✓ 884ms | ✓ 1591ms | ✓ 1252ms | 否 | ✓ 1610ms | http |
| 45.167.124.71:999 | ✓ 1863ms | ✓ 1545ms | ✓ 1460ms | ✓ 1709ms | ✓ 1402ms | http |
| 72.11.150.178:6005 | ✓ 452ms | ✓ 1163ms | ✓ 776ms | ✓ 1272ms | ✓ 1001ms | http |
| 206.206.126.177:2412 | ✓ 1127ms | 否 | ✓ 1444ms | ✓ 1215ms | ✓ 960ms | http |
| 109.120.156.122:8090 | ✓ 1128ms | ✓ 1955ms | ✓ 1563ms | 否 | ✓ 1712ms | http |
| 148.230.4.241:999 | ✓ 1050ms | ✓ 1912ms | ✓ 652ms | ✓ 1513ms | ✓ 1503ms | http |
| 180.180.109.189:8080 | ✓ 1705ms | 否 | ✓ 1927ms | ✓ 1998ms | 否 | http |
| 119.195.17.15:3176 | ✓ 1866ms | ✓ 1217ms | ✓ 1060ms | ✓ 1482ms | ✓ 1365ms | http |
| 103.35.190.69:1081 | ✓ 139ms | ✓ 936ms | ✓ 84ms | ✓ 1389ms | ✓ 750ms | http |
| 223.206.59.59:8080 | 否 | 否 | ✓ 1514ms | ✓ 1740ms | ✓ 1727ms | http |
| 80.92.204.47:1081 | ✓ 1956ms | ✓ 1180ms | ✓ 911ms | ✓ 1421ms | 否 | http |
| 47.85.51.197:1080 | ✓ 861ms | 否 | ✓ 392ms | ✓ 1864ms | 否 | http |
| 45.59.122.132:80 | ✓ 1510ms | ✓ 1711ms | ✓ 1892ms | 否 | 否 | http |
| 47.77.216.82:1080 | ✓ 823ms | ✓ 1268ms | ✓ 849ms | ✓ 1043ms | ✓ 840ms | http |
| 149.51.42.10:8080 | ✓ 873ms | ✓ 1199ms | 否 | ✓ 1242ms | 否 | http |
| 149.51.42.10:3128 | ✓ 873ms | ✓ 1208ms | 否 | ✓ 1248ms | 否 | http |
| 129.159.159.78:3128 | ✓ 1260ms | ✓ 1846ms | 否 | 否 | ✓ 1908ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1768ms | ✓ 1953ms | ✓ 1384ms | http |
| 103.157.200.126:3128 | ✓ 1183ms | 否 | ✓ 1203ms | ✓ 1669ms | ✓ 1439ms | http |
| 86.104.72.219:1081 | ✓ 542ms | ✓ 1061ms | ✓ 591ms | ✓ 1271ms | 否 | http |
| 107.173.42.121:7890 | 否 | ✓ 1016ms | ✓ 808ms | ✓ 987ms | 否 | http |
| 86.104.72.219:1082 | 否 | 否 | ✓ 157ms | ✓ 1087ms | ✓ 1856ms | http |
| 34.96.238.40:8080 | ✓ 1324ms | 否 | ✓ 1072ms | 否 | ✓ 1271ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1715ms | ✓ 1804ms | ✓ 1485ms | http |
| 8.211.166.184:8081 | ✓ 1653ms | ✓ 1003ms | ✓ 840ms | ✓ 1065ms | ✓ 829ms | http |
| 1.231.81.166:3128 | ✓ 1674ms | ✓ 1163ms | ✓ 1758ms | ✓ 1261ms | ✓ 1038ms | http |
| 152.70.91.193:40000 | ✓ 1895ms | 否 | 否 | ✓ 1928ms | ✓ 1859ms | http |
| 103.35.190.69:1082 | ✓ 149ms | ✓ 1644ms | ✓ 1894ms | 否 | ✓ 1035ms | http |
| 212.58.132.5:8888 | ✓ 1739ms | 否 | ✓ 1478ms | ✓ 1448ms | ✓ 1188ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1490ms | ✓ 1199ms | ✓ 1571ms | 否 | http |
| 8.219.97.248:80 | ✓ 1414ms | 否 | ✓ 1367ms | ✓ 1710ms | 否 | http |
| 43.133.44.89:8888 | ✓ 904ms | 否 | ✓ 1712ms | ✓ 1208ms | 否 | http |
| 116.171.106.26:3443 | ✓ 1790ms | 否 | ✓ 1648ms | ✓ 1965ms | 否 | http |
| 103.247.242.22:8080 | 否 | 否 | ✓ 1590ms | ✓ 1747ms | ✓ 1440ms | http |
| 154.27.196.3:8080 | ✓ 1000ms | ✓ 1517ms | ✓ 1209ms | ✓ 1732ms | ✓ 1480ms | http |
| 105.159.150.118:4029 | ✓ 1155ms | 否 | ✓ 672ms | 否 | ✓ 1359ms | http |
| 204.157.251.234:999 | ✓ 1373ms | ✓ 1493ms | ✓ 1181ms | ✓ 1577ms | ✓ 1561ms | http |
| 105.159.127.88:4029 | ✓ 1160ms | ✓ 1989ms | ✓ 759ms | 否 | ✓ 1345ms | http |
| 108.181.0.167:8080 | ✓ 577ms | ✓ 1125ms | ✓ 1273ms | ✓ 966ms | ✓ 907ms | http |
| 129.213.162.27:17777 | ✓ 1107ms | ✓ 1585ms | 否 | ✓ 1884ms | 否 | http |
| 3.101.133.120:80 | ✓ 432ms | ✓ 1443ms | ✓ 1616ms | ✓ 1556ms | ✓ 1095ms | http |
| 144.91.102.48:3128 | ✓ 912ms | ✓ 1468ms | ✓ 1577ms | ✓ 1557ms | ✓ 1196ms | http |
| 118.113.246.229:1080 | ✓ 1392ms | ✓ 1839ms | ✓ 1576ms | 否 | ✓ 1487ms | http |
| 104.248.81.109:3128 | 否 | ✓ 1589ms | ✓ 1780ms | 否 | ✓ 1557ms | http |
| 152.42.170.187:9090 | ✓ 1092ms | 否 | ✓ 1083ms | ✓ 1695ms | 否 | http |
| 86.104.72.220:1081 | ✓ 587ms | ✓ 923ms | ✓ 85ms | 否 | 否 | http |
| 86.104.72.220:1082 | ✓ 509ms | ✓ 878ms | ✓ 819ms | ✓ 1053ms | ✓ 722ms | http |
| 45.140.147.82:1081 | ✓ 1542ms | 否 | ✓ 1884ms | 否 | ✓ 1397ms | http |
| 45.140.147.82:1082 | ✓ 1546ms | 否 | ✓ 1885ms | 否 | ✓ 1420ms | http |
| 20.164.75.153:8080 | ✓ 1532ms | 否 | ✓ 1158ms | 否 | ✓ 1761ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1185ms | 否 | ✓ 1915ms | ✓ 1178ms | http |
| 103.35.191.244:1082 | 否 | ✓ 1646ms | ✓ 66ms | ✓ 1234ms | ✓ 1313ms | http |
| 62.113.119.14:8080 | ✓ 574ms | 否 | ✓ 1854ms | ✓ 1578ms | 否 | http |
| 117.236.124.166:3128 | ✓ 961ms | 否 | ✓ 1043ms | 否 | ✓ 1069ms | http |
| 202.129.206.239:3128 | ✓ 1554ms | 否 | 否 | ✓ 1525ms | ✓ 1721ms | http |
| 101.32.243.189:80 | ✓ 1530ms | ✓ 1703ms | ✓ 1626ms | ✓ 1643ms | ✓ 1527ms | http |
| 61.52.131.172:8443 | ✓ 936ms | ✓ 1135ms | ✓ 1013ms | ✓ 1147ms | ✓ 927ms | http |
| 218.108.131.186:17890 | ✓ 886ms | ✓ 1086ms | ✓ 901ms | ✓ 1196ms | ✓ 895ms | http |
| 60.249.94.208:3128 | 否 | ✓ 1208ms | ✓ 1114ms | ✓ 1236ms | ✓ 1010ms | http |

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
