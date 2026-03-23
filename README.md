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

最后更新：2026-03-23 12:42:12 UTC（2026-03-23 20:42:12 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1496ms | 否 | ✓ 1011ms | ✓ 1023ms | ✓ 1106ms | http |
| 113.160.132.26:8080 | ✓ 1871ms | ✓ 1669ms | ✓ 1271ms | 否 | ✓ 1119ms | http |
| 45.167.124.52:8080 | ✓ 1037ms | 否 | ✓ 1497ms | 否 | ✓ 1514ms | http |
| 167.103.34.108:8800 | ✓ 1288ms | 否 | ✓ 1374ms | ✓ 1456ms | ✓ 1383ms | http |
| 120.92.212.16:7890 | ✓ 1064ms | ✓ 1475ms | ✓ 1403ms | ✓ 1804ms | 否 | http |
| 219.117.204.211:7799 | ✓ 1500ms | 否 | ✓ 846ms | ✓ 985ms | ✓ 1268ms | http |
| 103.166.185.54:3128 | ✓ 1778ms | 否 | ✓ 1243ms | ✓ 1346ms | ✓ 1209ms | http |
| 167.103.31.122:8800 | 否 | 否 | ✓ 1292ms | ✓ 1629ms | ✓ 1552ms | http |
| 116.80.49.168:3172 | ✓ 1936ms | 否 | ✓ 1605ms | ✓ 1963ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1804ms | 否 | ✓ 1726ms | ✓ 1414ms | 否 | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 1018ms | ✓ 1002ms | ✓ 763ms | http |
| 147.161.239.240:8800 | ✓ 572ms | ✓ 1618ms | ✓ 1401ms | ✓ 1642ms | ✓ 1546ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1284ms | ✓ 1149ms | ✓ 1236ms | 否 | http |
| 38.145.218.217:8444 | ✓ 565ms | ✓ 960ms | ✓ 1452ms | ✓ 1704ms | ✓ 1766ms | http |
| 47.95.231.180:8084 | ✓ 807ms | ✓ 1166ms | ✓ 915ms | ✓ 1114ms | ✓ 872ms | http |
| 62.113.119.14:8080 | ✓ 640ms | ✓ 1584ms | ✓ 581ms | ✓ 1615ms | ✓ 1111ms | http |
| 144.31.79.117:8888 | ✓ 712ms | 否 | ✓ 969ms | ✓ 1824ms | ✓ 1416ms | http |
| 121.230.8.14:1080 | ✓ 1193ms | ✓ 1541ms | 否 | ✓ 1658ms | 否 | http |
| 201.144.20.238:3128 | ✓ 702ms | ✓ 1327ms | ✓ 1590ms | ✓ 1541ms | ✓ 1494ms | http |
| 185.115.74.185:8080 | ✓ 879ms | ✓ 1721ms | ✓ 1744ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 1470ms | 否 | ✓ 1219ms | 否 | ✓ 1777ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 961ms | ✓ 994ms | ✓ 1007ms | http |
| 137.220.150.22:6005 | ✓ 1126ms | 否 | 否 | ✓ 1403ms | ✓ 1086ms | http |
| 38.34.179.39:8452 | ✓ 345ms | ✓ 944ms | ✓ 521ms | ✓ 1119ms | ✓ 1298ms | http |
| 104.129.202.127:10810 | ✓ 254ms | 否 | ✓ 761ms | ✓ 1009ms | ✓ 774ms | http |
| 1.231.81.166:3128 | ✓ 1253ms | ✓ 1913ms | ✓ 1997ms | ✓ 1285ms | ✓ 1355ms | http |
| 103.84.95.54:7890 | ✓ 1653ms | 否 | ✓ 1439ms | 否 | ✓ 1864ms | http |
| 38.145.208.242:8451 | ✓ 903ms | 否 | ✓ 1126ms | ✓ 1034ms | ✓ 1226ms | http |
| 38.180.2.107:3128 | ✓ 1078ms | ✓ 1943ms | ✓ 1643ms | 否 | ✓ 1672ms | http |
| 88.80.150.82:8080 | ✓ 962ms | 否 | 否 | ✓ 1976ms | ✓ 1710ms | https |
| 112.111.13.253:7890 | ✓ 929ms | ✓ 1112ms | ✓ 1104ms | ✓ 1293ms | ✓ 902ms | http |
| 59.46.216.131:30001 | ✓ 1192ms | 否 | 否 | ✓ 1743ms | ✓ 1129ms | http |
| 45.136.198.40:3128 | ✓ 1057ms | 否 | ✓ 1765ms | 否 | ✓ 1658ms | http |
| 166.249.54.61:7234 | ✓ 1427ms | ✓ 1938ms | 否 | ✓ 1931ms | ✓ 1755ms | http |
| 147.161.246.37:10801 | ✓ 865ms | 否 | ✓ 1360ms | ✓ 1981ms | ✓ 1750ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 382ms | ✓ 1563ms | ✓ 1163ms | http |
| 172.212.68.37:3128 | ✓ 1036ms | 否 | ✓ 1333ms | ✓ 1238ms | ✓ 1026ms | http |
| 38.145.208.172:8448 | 否 | 否 | ✓ 276ms | ✓ 895ms | ✓ 1069ms | http |
| 58.220.95.8:10174 | ✓ 977ms | ✓ 1222ms | ✓ 1535ms | ✓ 1246ms | ✓ 1267ms | http |
| 166.88.55.83:7890 | ✓ 757ms | ✓ 1625ms | ✓ 804ms | ✓ 934ms | ✓ 745ms | http |
| 202.141.161.53:30001 | 否 | 否 | ✓ 1204ms | ✓ 1340ms | ✓ 1208ms | http |
| 210.45.70.16:7895 | ✓ 1206ms | 否 | ✓ 1548ms | ✓ 1434ms | ✓ 1139ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1463ms | 否 | ✓ 1497ms | ✓ 1966ms | http |
| 150.249.255.91:3128 | ✓ 727ms | 否 | ✓ 682ms | ✓ 986ms | ✓ 1340ms | http |
| 46.101.190.71:3128 | ✓ 458ms | ✓ 1601ms | ✓ 1495ms | ✓ 1965ms | ✓ 1627ms | http |
| 101.47.73.135:3128 | ✓ 1411ms | 否 | 否 | ✓ 1647ms | ✓ 1315ms | http |
| 223.16.170.103:80 | ✓ 1271ms | 否 | ✓ 1217ms | ✓ 1555ms | 否 | http |
| 160.250.4.245:1 | ✓ 1513ms | 否 | ✓ 1525ms | ✓ 1490ms | 否 | http |
| 8.212.130.232:8080 | ✓ 1074ms | ✓ 1936ms | ✓ 1033ms | 否 | 否 | http |
| 64.227.76.27:1080 | ✓ 1807ms | 否 | ✓ 1018ms | 否 | ✓ 1259ms | http |
| 195.123.209.48:3128 | ✓ 983ms | ✓ 1405ms | ✓ 1768ms | 否 | ✓ 1636ms | http |
| 139.159.99.242:8080 | ✓ 1755ms | ✓ 1218ms | ✓ 1342ms | 否 | 否 | http |
| 91.236.238.103:8080 | ✓ 1287ms | 否 | ✓ 913ms | ✓ 1609ms | ✓ 1545ms | http |
| 103.183.10.169:3125 | ✓ 1910ms | 否 | ✓ 1808ms | ✓ 1847ms | ✓ 1559ms | http |
| 59.8.203.55:80 | ✓ 1670ms | ✓ 1571ms | ✓ 1363ms | ✓ 1693ms | ✓ 1125ms | http |
| 128.199.121.61:9090 | ✓ 1718ms | 否 | 否 | ✓ 1441ms | ✓ 1674ms | http |

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
