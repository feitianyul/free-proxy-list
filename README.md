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

最后更新：2026-03-15 06:03:05 UTC（2026-03-15 14:03:05 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 677ms | 否 | ✓ 1429ms | ✓ 1170ms | ✓ 877ms | http |
| 113.160.132.26:8080 | ✓ 1711ms | 否 | 否 | ✓ 1167ms | ✓ 912ms | http |
| 45.167.124.52:8080 | ✓ 1285ms | 否 | ✓ 1396ms | ✓ 1689ms | ✓ 1361ms | http |
| 120.92.212.16:7890 | ✓ 922ms | ✓ 1234ms | ✓ 1249ms | ✓ 1241ms | ✓ 971ms | http |
| 45.136.198.40:3128 | ✓ 766ms | ✓ 1931ms | 否 | 否 | ✓ 1839ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1864ms | ✓ 1497ms | ✓ 951ms | http |
| 35.225.22.61:80 | 否 | ✓ 1364ms | ✓ 236ms | ✓ 1377ms | 否 | http |
| 120.92.212.16:8890 | ✓ 935ms | ✓ 1226ms | ✓ 1210ms | ✓ 1219ms | ✓ 972ms | http |
| 62.60.177.204:34094 | ✓ 674ms | 否 | 否 | ✓ 1126ms | ✓ 996ms | http |
| 101.43.127.100:8877 | ✓ 959ms | ✓ 1045ms | ✓ 1099ms | ✓ 1198ms | ✓ 868ms | http |
| 81.70.169.194:80 | ✓ 1238ms | ✓ 1301ms | ✓ 1112ms | ✓ 1358ms | ✓ 969ms | http |
| 210.223.44.230:3128 | ✓ 625ms | ✓ 1216ms | ✓ 1383ms | ✓ 1942ms | 否 | http |
| 85.198.96.242:3128 | ✓ 959ms | 否 | 否 | ✓ 1987ms | ✓ 1338ms | http |
| 101.43.255.96:80 | ✓ 1238ms | ✓ 1329ms | ✓ 1208ms | 否 | ✓ 987ms | http |
| 160.250.4.245:1 | ✓ 1015ms | 否 | ✓ 1366ms | ✓ 1231ms | ✓ 1001ms | http |
| 160.250.4.13:1 | ✓ 1014ms | 否 | ✓ 1303ms | ✓ 1279ms | ✓ 1003ms | http |
| 86.53.183.16:1080 | ✓ 863ms | 否 | ✓ 1665ms | 否 | ✓ 1681ms | http |
| 62.234.206.73:3128 | ✓ 1201ms | 否 | ✓ 1225ms | 否 | ✓ 1479ms | http |
| 38.145.203.135:8443 | ✓ 611ms | 否 | ✓ 321ms | ✓ 741ms | ✓ 1611ms | http |
| 45.136.130.232:8443 | 否 | 否 | ✓ 1947ms | ✓ 1745ms | ✓ 686ms | http |
| 45.136.130.233:8443 | 否 | 否 | ✓ 1945ms | ✓ 1767ms | ✓ 665ms | http |
| 38.145.203.235:8443 | 否 | 否 | ✓ 1945ms | ✓ 1752ms | ✓ 1018ms | http |
| 144.124.227.90:21074 | ✓ 632ms | 否 | ✓ 1829ms | ✓ 1869ms | 否 | http |
| 38.55.106.208:6005 | ✓ 1610ms | 否 | ✓ 1003ms | ✓ 924ms | ✓ 679ms | http |
| 212.192.13.76:6005 | ✓ 1645ms | 否 | ✓ 1314ms | ✓ 1169ms | ✓ 1322ms | http |
| 64.188.90.36:1080 | ✓ 1264ms | 否 | ✓ 854ms | ✓ 1967ms | 否 | http |
| 146.19.128.135:1080 | ✓ 1261ms | 否 | ✓ 858ms | ✓ 1958ms | 否 | http |
| 165.227.5.10:8888 | 否 | ✓ 950ms | ✓ 362ms | 否 | ✓ 708ms | http |
| 8.219.97.248:80 | ✓ 1199ms | 否 | ✓ 1506ms | 否 | ✓ 1241ms | http |
| 101.47.73.135:3128 | ✓ 966ms | 否 | ✓ 1923ms | ✓ 1556ms | 否 | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1335ms | ✓ 1940ms | ✓ 1291ms | http |
| 95.3.9.78:3128 | ✓ 1695ms | 否 | ✓ 1474ms | 否 | ✓ 1424ms | http |
| 34.96.238.40:8080 | ✓ 1204ms | 否 | 否 | ✓ 1035ms | ✓ 1067ms | http |
| 149.62.191.202:3128 | ✓ 826ms | 否 | ✓ 1753ms | ✓ 1983ms | ✓ 1470ms | http |
| 95.3.9.78:8080 | ✓ 923ms | 否 | ✓ 1207ms | 否 | ✓ 1531ms | http |
| 24.199.124.152:3128 | ✓ 1848ms | ✓ 903ms | ✓ 1186ms | ✓ 742ms | ✓ 607ms | http |
| 45.22.209.157:8888 | ✓ 1447ms | 否 | ✓ 413ms | ✓ 1773ms | 否 | http |
| 106.117.208.101:7890 | ✓ 984ms | ✓ 1257ms | ✓ 1361ms | ✓ 1336ms | ✓ 1056ms | http |
| 137.220.150.152:6005 | ✓ 749ms | 否 | ✓ 800ms | ✓ 1098ms | ✓ 906ms | http |
| 137.220.151.110:6005 | ✓ 852ms | 否 | ✓ 795ms | ✓ 1076ms | ✓ 853ms | http |
| 137.220.150.104:6005 | ✓ 833ms | 否 | ✓ 761ms | ✓ 1112ms | ✓ 897ms | http |
| 38.55.105.94:6005 | ✓ 870ms | 否 | ✓ 975ms | ✓ 1207ms | ✓ 1056ms | http |
| 38.55.106.245:6005 | ✓ 1026ms | 否 | ✓ 1993ms | ✓ 926ms | ✓ 994ms | http |
| 92.119.127.213:6005 | ✓ 1058ms | 否 | ✓ 1615ms | ✓ 1717ms | ✓ 1914ms | http |
| 38.145.208.217:8443 | ✓ 710ms | ✓ 1098ms | ✓ 899ms | 否 | 否 | http |
| 38.145.208.223:8443 | ✓ 192ms | ✓ 1250ms | ✓ 144ms | 否 | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1175ms | ✓ 1138ms | ✓ 1322ms | ✓ 1254ms | http |
| 159.223.42.219:3128 | ✓ 823ms | 否 | ✓ 1172ms | ✓ 1055ms | ✓ 866ms | http |
| 59.46.216.131:30001 | ✓ 972ms | 否 | ✓ 1102ms | 否 | ✓ 1023ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1442ms | ✓ 1339ms | ✓ 983ms | http |
| 72.11.150.178:6005 | ✓ 1121ms | ✓ 1284ms | ✓ 1064ms | ✓ 1171ms | ✓ 1102ms | http |
| 38.207.164.82:6005 | ✓ 1666ms | 否 | ✓ 1058ms | ✓ 1246ms | ✓ 1022ms | http |
| 103.113.70.189:1081 | ✓ 313ms | 否 | 否 | ✓ 1142ms | ✓ 846ms | http |
| 103.39.51.190:8080 | ✓ 1639ms | 否 | 否 | ✓ 1568ms | ✓ 1330ms | http |
| 61.52.131.172:8443 | ✓ 865ms | ✓ 1093ms | ✓ 1118ms | ✓ 1218ms | ✓ 894ms | http |
| 47.101.149.27:9010 | ✓ 1706ms | ✓ 1296ms | 否 | 否 | ✓ 1351ms | http |

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
