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

最后更新：2026-04-25 22:31:26 UTC（2026-04-26 06:31:26 UTC+8）

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
| 206.206.126.177:2412 | ✓ 1495ms | ✓ 1451ms | ✓ 1323ms | ✓ 1192ms | ✓ 1086ms | http |
| 47.85.51.197:1080 | ✓ 446ms | ✓ 1017ms | ✓ 1052ms | 否 | ✓ 1204ms | http |
| 62.113.119.14:8080 | ✓ 1686ms | ✓ 1950ms | ✓ 763ms | ✓ 1501ms | 否 | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1523ms | ✓ 1639ms | ✓ 1449ms | http |
| 80.92.204.47:1081 | ✓ 683ms | 否 | 否 | ✓ 1757ms | ✓ 1242ms | http |
| 42.200.76.16:3888 | ✓ 895ms | 否 | ✓ 910ms | ✓ 1682ms | 否 | http |
| 90.174.128.42:3128 | 否 | 否 | ✓ 1772ms | ✓ 1949ms | ✓ 1599ms | http |
| 120.92.108.86:7890 | ✓ 1345ms | 否 | ✓ 1816ms | 否 | ✓ 1433ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1328ms | ✓ 1388ms | ✓ 1574ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1160ms | ✓ 1774ms | ✓ 1232ms | 否 | ✓ 1551ms | http |
| 36.141.21.200:7890 | ✓ 1049ms | ✓ 1324ms | ✓ 1134ms | ✓ 1379ms | ✓ 1096ms | http |
| 47.84.76.30:1080 | ✓ 1687ms | 否 | ✓ 1082ms | ✓ 1234ms | ✓ 961ms | http |
| 152.32.132.190:7890 | ✓ 1045ms | 否 | ✓ 1662ms | 否 | ✓ 853ms | http |
| 177.93.132.244:3128 | ✓ 879ms | 否 | ✓ 1038ms | 否 | ✓ 1748ms | http |
| 121.230.8.41:1080 | ✓ 1111ms | ✓ 1566ms | ✓ 1361ms | ✓ 1553ms | ✓ 1182ms | http |
| 64.181.240.152:3128 | ✓ 1161ms | ✓ 1969ms | ✓ 1698ms | ✓ 984ms | ✓ 811ms | http |
| 47.84.59.16:1080 | ✓ 1079ms | ✓ 1782ms | ✓ 984ms | ✓ 1243ms | ✓ 986ms | http |
| 82.148.18.242:443 | ✓ 686ms | ✓ 1773ms | ✓ 1493ms | 否 | ✓ 1682ms | http |
| 47.105.98.23:3128 | 否 | 否 | ✓ 1319ms | ✓ 1362ms | ✓ 1368ms | http |
| 59.46.216.131:30001 | ✓ 1092ms | ✓ 1481ms | ✓ 1087ms | ✓ 1486ms | 否 | http |
| 94.131.122.129:1081 | ✓ 1579ms | ✓ 1525ms | ✓ 917ms | ✓ 1514ms | ✓ 1228ms | http |
| 218.108.131.186:17890 | ✓ 1220ms | ✓ 1654ms | ✓ 1596ms | 否 | 否 | http |
| 86.104.74.110:1081 | ✓ 1654ms | ✓ 1437ms | ✓ 1599ms | 否 | ✓ 1636ms | http |
| 117.236.124.166:3128 | ✓ 1323ms | 否 | ✓ 1509ms | ✓ 1966ms | ✓ 1525ms | http |
| 84.47.150.125:1080 | ✓ 696ms | ✓ 1477ms | 否 | 否 | ✓ 1861ms | http |
| 2.27.54.161:1080 | ✓ 729ms | ✓ 1876ms | ✓ 740ms | ✓ 1757ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1879ms | ✓ 1830ms | ✓ 1198ms | 否 | 否 | http |
| 47.112.25.109:7890 | 否 | ✓ 1698ms | ✓ 1184ms | ✓ 1390ms | 否 | http |
| 159.89.31.62:8080 | ✓ 502ms | ✓ 1898ms | 否 | ✓ 1944ms | 否 | http |
| 126.209.18.142:8082 | 否 | 否 | ✓ 1734ms | ✓ 1608ms | ✓ 1624ms | http |
| 45.140.147.155:1081 | ✓ 1783ms | ✓ 1440ms | ✓ 862ms | 否 | ✓ 1205ms | http |
| 45.140.147.155:1082 | ✓ 659ms | ✓ 1913ms | ✓ 1503ms | 否 | ✓ 1198ms | http |
| 43.133.90.161:8888 | ✓ 1986ms | ✓ 1427ms | ✓ 1700ms | 否 | 否 | http |
| 183.232.248.73:7890 | ✓ 1329ms | ✓ 1465ms | ✓ 1401ms | ✓ 1831ms | ✓ 1365ms | http |
| 8.211.166.184:8081 | ✓ 1645ms | ✓ 1284ms | ✓ 854ms | ✓ 1030ms | ✓ 812ms | http |
| 20.120.225.109:3128 | ✓ 440ms | ✓ 1259ms | ✓ 983ms | ✓ 1268ms | ✓ 1122ms | http |
| 47.101.159.19:8899 | ✓ 1003ms | ✓ 1206ms | ✓ 1077ms | ✓ 1262ms | ✓ 1006ms | http |
| 45.76.207.177:40000 | ✓ 1634ms | 否 | ✓ 994ms | ✓ 1510ms | ✓ 1680ms | http |
| 94.131.122.128:1081 | ✓ 1174ms | ✓ 1449ms | ✓ 662ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1846ms | 否 | 否 | ✓ 1325ms | ✓ 1881ms | http |
| 47.84.73.61:1080 | ✓ 1243ms | ✓ 1906ms | ✓ 1055ms | ✓ 1247ms | ✓ 1011ms | http |
| 60.249.94.208:3128 | ✓ 1093ms | ✓ 1371ms | ✓ 1089ms | ✓ 1220ms | 否 | http |
| 68.183.199.89:1080 | ✓ 1356ms | 否 | ✓ 356ms | ✓ 930ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1408ms | ✓ 1927ms | ✓ 1546ms | ✓ 1393ms | ✓ 878ms | http |
| 8.219.195.129:1080 | ✓ 982ms | ✓ 1840ms | ✓ 1083ms | ✓ 1253ms | ✓ 983ms | http |
| 180.103.19.143:1080 | ✓ 1034ms | ✓ 1280ms | ✓ 1205ms | ✓ 1571ms | ✓ 1187ms | http |
| 121.230.8.144:1080 | ✓ 1596ms | ✓ 1609ms | ✓ 1245ms | ✓ 1431ms | ✓ 1267ms | http |
| 31.131.248.48:3129 | ✓ 933ms | ✓ 1712ms | ✓ 1317ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1648ms | 否 | ✓ 1918ms | ✓ 1954ms | ✓ 1603ms | http |
| 82.114.228.67:1080 | 否 | ✓ 1926ms | ✓ 1309ms | ✓ 1543ms | 否 | http |

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
