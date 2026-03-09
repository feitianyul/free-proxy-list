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

最后更新：2026-03-09 16:57:33 UTC（2026-03-10 00:57:33 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1660ms | ✓ 1353ms | ✓ 992ms | ✓ 1106ms | ✓ 802ms | http |
| 61.72.110.114:3128 | ✓ 1668ms | 否 | ✓ 1863ms | ✓ 1774ms | ✓ 897ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1950ms | 否 | ✓ 1564ms | ✓ 1081ms | http |
| 165.227.5.10:8888 | ✓ 1393ms | ✓ 1009ms | ✓ 1103ms | ✓ 1113ms | ✓ 1731ms | http |
| 115.231.181.40:8128 | ✓ 1268ms | ✓ 1246ms | 否 | 否 | ✓ 1231ms | http |
| 152.42.213.210:8080 | ✓ 1508ms | 否 | 否 | ✓ 1412ms | ✓ 932ms | http |
| 35.225.22.61:80 | ✓ 462ms | ✓ 1746ms | 否 | ✓ 1050ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1247ms | ✓ 1933ms | ✓ 1265ms | 否 | 否 | http |
| 121.237.181.137:8888 | ✓ 928ms | ✓ 1174ms | ✓ 981ms | ✓ 1314ms | ✓ 1985ms | http |
| 81.70.169.194:80 | ✓ 1097ms | 否 | ✓ 1287ms | ✓ 1406ms | ✓ 1729ms | http |
| 101.43.255.96:80 | ✓ 1980ms | 否 | ✓ 1475ms | ✓ 1592ms | 否 | http |
| 190.9.109.198:999 | ✓ 1843ms | ✓ 1242ms | ✓ 1630ms | ✓ 1377ms | ✓ 1307ms | http |
| 190.9.109.207:999 | ✓ 1208ms | ✓ 1557ms | ✓ 1481ms | 否 | ✓ 1177ms | http |
| 194.213.18.200:443 | ✓ 1810ms | ✓ 1712ms | ✓ 811ms | ✓ 1423ms | ✓ 875ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1191ms | ✓ 1100ms | ✓ 1199ms | http |
| 101.47.73.135:3128 | ✓ 1384ms | 否 | ✓ 1222ms | 否 | ✓ 1192ms | http |
| 152.42.213.210:80 | ✓ 804ms | 否 | ✓ 1145ms | ✓ 1147ms | ✓ 915ms | http |
| 125.128.12.14:3128 | ✓ 1985ms | ✓ 1594ms | ✓ 886ms | ✓ 1430ms | ✓ 908ms | http |
| 116.80.49.161:3172 | ✓ 1693ms | 否 | ✓ 1620ms | 否 | ✓ 1816ms | http |
| 67.169.98.211:443 | ✓ 996ms | 否 | ✓ 519ms | ✓ 1113ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1329ms | ✓ 1289ms | 否 | ✓ 1051ms | http |
| 168.235.110.63:3128 | ✓ 311ms | ✓ 1495ms | ✓ 1137ms | ✓ 1109ms | ✓ 872ms | http |
| 14.56.107.244:3128 | ✓ 1820ms | ✓ 934ms | ✓ 646ms | ✓ 1078ms | ✓ 883ms | http |
| 61.72.221.94:3128 | ✓ 918ms | ✓ 1255ms | ✓ 1122ms | ✓ 1109ms | ✓ 1092ms | http |
| 125.128.12.144:3128 | ✓ 928ms | 否 | 否 | ✓ 1462ms | ✓ 860ms | http |
| 8.219.97.248:80 | ✓ 1437ms | 否 | ✓ 1518ms | ✓ 1497ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1615ms | ✓ 1207ms | ✓ 1012ms | ✓ 1340ms | ✓ 1241ms | http |
| 88.80.150.82:8080 | ✓ 1698ms | ✓ 1880ms | 否 | 否 | ✓ 1865ms | https |
| 205.209.118.30:3138 | ✓ 1663ms | 否 | ✓ 1938ms | ✓ 1198ms | ✓ 978ms | http |
| 120.55.163.237:10086 | ✓ 920ms | ✓ 1139ms | ✓ 1041ms | ✓ 1175ms | ✓ 978ms | http |
| 45.140.147.155:1082 | ✓ 543ms | 否 | ✓ 1302ms | ✓ 1482ms | ✓ 1152ms | http |
| 45.140.147.155:1081 | ✓ 556ms | 否 | ✓ 1210ms | ✓ 1596ms | ✓ 1093ms | http |
| 47.110.42.192:9003 | ✓ 1776ms | ✓ 1692ms | ✓ 1418ms | ✓ 1976ms | 否 | http |
| 43.167.227.161:1080 | 否 | ✓ 1585ms | ✓ 972ms | ✓ 1441ms | 否 | http |
| 113.176.92.71:3128 | 否 | ✓ 1730ms | ✓ 1310ms | ✓ 1271ms | ✓ 1008ms | http |
| 178.236.245.59:3128 | ✓ 1090ms | 否 | ✓ 1255ms | 否 | ✓ 1981ms | http |
| 45.140.147.82:1081 | ✓ 1017ms | 否 | ✓ 1200ms | ✓ 1536ms | 否 | http |
| 14.225.212.37:7890 | 否 | ✓ 1450ms | ✓ 874ms | ✓ 1262ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1097ms | ✓ 1728ms | 否 | 否 | ✓ 1217ms | http |
| 210.77.29.245:7890 | ✓ 1001ms | ✓ 1174ms | 否 | ✓ 1248ms | ✓ 999ms | http |
| 121.230.8.208:1080 | ✓ 1011ms | ✓ 1346ms | ✓ 1541ms | ✓ 1304ms | ✓ 1297ms | http |
| 121.230.8.213:1080 | ✓ 1197ms | ✓ 1488ms | ✓ 1173ms | ✓ 1518ms | ✓ 1173ms | http |
| 121.230.8.80:1080 | ✓ 1282ms | ✓ 1511ms | 否 | ✓ 1766ms | ✓ 1226ms | http |
| 103.39.51.190:8080 | ✓ 1762ms | 否 | 否 | ✓ 1887ms | ✓ 1857ms | http |
| 39.104.201.40:7890 | ✓ 973ms | ✓ 1285ms | ✓ 1069ms | ✓ 1283ms | ✓ 1030ms | http |
| 5.129.206.247:8888 | ✓ 1318ms | 否 | ✓ 1730ms | 否 | ✓ 1931ms | http |
| 136.49.34.18:8888 | ✓ 816ms | ✓ 1487ms | ✓ 1605ms | ✓ 1310ms | ✓ 1202ms | http |

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
